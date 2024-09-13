import * as core from "@actions/core";
import jira from "jira.js";
import axios from "axios";
import { join } from "path";
import { createJiraClient, extractJiraIssueNumbersFrom, getJiraEnvVars, PR_PREFIX, SOLIDITY_REVIEW_PREFIX } from "./lib";
import { appendIssueNumberToChangesetFile, extractChangesetFiles } from "./changeset-lib";

async function main() {
    core.info('Started linking PR to a Solidity Review issue')
    const solidityReviewTemplateKey = readSolidityReviewTemplateKey()
    const changesetFiles = extractChangesetFiles();

    if (changesetFiles.length > 1) {
      core.setFailed(
        `Solidity Review enforcement only works with 1 changeset per PR, but found ${changesetFiles.length} changesets`
      );
      return
    }

    const changesetFile = changesetFiles[0]

    // first let's make sure that this PR is linked to a JIRA issue, we will need it later anyway
    const jiraPRIssues = await extractJiraIssueNumbersFrom(PR_PREFIX, changesetFiles)
    if (jiraPRIssues.length !== 1) {
        core.setFailed(
            `Solidity Review enforcement only works with 1 JIRA issue per PR, but found ${jiraPRIssues.length} issues in changeset file ${changesetFile}`
          );
          return
    }

    const jiraPRIssue = jiraPRIssues[0]

    // now let's check whether the issue is already linked to at least one Solidity Review issue (it's okay if there's more than one if PR modifies files for more than one project)
    const jiraSolidityIssues = await extractJiraIssueNumbersFrom(SOLIDITY_REVIEW_PREFIX, changesetFiles)
    if (jiraSolidityIssues.length > 0) {
        core.info(`Found linked Solidity Review issue(s): ${join(...jiraSolidityIssues)}. Nothing more needs to be done.`)

        return
    }

    const jiraProject = extracProjectFromIssueKey(jiraPRIssue)
    if (!jiraProject) {
        core.setFailed(`Could not extract project key from issue: ${jiraPRIssue}`)

        return
    }

    const client = createJiraClient();

    core.info(`Checking if project ${jiraProject} has any open Solidity Review issues`)
    let solidityReviewIssueKey: string
    const openSolidityReviewIssues = await getOpenSolidityReviewIssuesForProject(client, jiraProject, [solidityReviewTemplateKey])
    if (openSolidityReviewIssues.length === 0) {
        solidityReviewIssueKey = await createSolidityReviewIssue(client, jiraProject, solidityReviewTemplateKey)
    } else if (openSolidityReviewIssues.length === 1) {
        solidityReviewIssueKey = openSolidityReviewIssues[0]
    } else {
        throw new Error(`Found following open Solidity Review issues for project ${jiraProject}: ${join(...openSolidityReviewIssues)}. Since we are unable to automatically determine, which one should be used, please manualy add it to changeset file: ${changesetFile}. Use this exact format:
${SOLIDITY_REVIEW_PREFIX}<issue-key>

Exmaple with issue key 'PROJ-1234':
${SOLIDITY_REVIEW_PREFIX}PROJ-1234`)
    }

    core.info(`Will use Solidity Review issue: ${solidityReviewIssueKey}`)
    await linkIssues(client, solidityReviewIssueKey, jiraPRIssue, 'Blocks')

    core.info(`Appending JIRA Solidity Review issue ${solidityReviewIssueKey} to changeset file`);
    await appendIssueNumberToChangesetFile(SOLIDITY_REVIEW_PREFIX, changesetFile, solidityReviewIssueKey);
    core.info('Finished linking PR to a Solidity Review issue')
  }

  async function getIssueKeys(client: jira.Version3Client, projectKey: string, issueType: string, title: string, status: string, issueKeysToIgnore: string[], maxResults?: number): Promise<string[]> {
    if (!maxResults) {
        maxResults = 10
    }
    try {
      let jql = `project = ${projectKey} AND issuetype = "${issueType}" AND summary ~ "${title}" AND status = "${status}"`;
      if (issueKeysToIgnore.length > 0) {
        jql = `${jql} AND issuekey NOT IN (${issueKeysToIgnore.join(',')})`
      }
      core.debug(`Searching for issue using jql: '${jql}'`)
      const result = await client.issueSearch.searchForIssuesUsingJql({
        jql: jql,
        maxResults: maxResults,
        fields: ['key']
      });

      if (result.issues === undefined) {
        core.debug('Found no matching issues.')
        return [];
      }

      return result.issues.map(issue => issue.key);
    } catch (error) {
      core.error('Error searching for issue: ' + error);
      return [];
    }
  }

  export async function linkIssues(client: jira.Version3Client, inwardIssueKey: string, outwardIssueKey: string, linkType: string) {
    core.debug(`Linking issue ${inwardIssueKey} to ${outwardIssueKey} with link type of '${linkType}'`)
    try {
      await client.issueLinks.linkIssues({
        type: {
          name: linkType,
        },
        inwardIssue: {
          key: inwardIssueKey,
        },
        outwardIssue: {
          key: outwardIssueKey,
        }
      });

      core.debug(`Successfully linked issues: ${inwardIssueKey} now '${linkType}' ${outwardIssueKey}`);
    } catch (error) {
        core.error(`Error linking issues ${inwardIssueKey} and ${outwardIssueKey}: ` + error);
        throw error
    }
  }

  async function createSolidityReviewIssue(client: jira.Version3Client, projectKey: string, sourceIssueKey: string) {
    let solidityReviewKey = ""
    try {
      core.info(`Creating new Solidity Review issue in project ${projectKey}`)
      solidityReviewKey = await cloneIssue(client, sourceIssueKey, projectKey)
      await cloneLinkedIssues(client, projectKey, sourceIssueKey, solidityReviewKey, ['Epic'], 2)
      core.info(`Created new Solidity Review issue in project ${projectKey}. Issue key: ${solidityReviewKey}`)

      return solidityReviewKey
    } catch (error) {
      core.setFailed('Failed to create new Solidity Review issue: ' + error)
      if (solidityReviewKey !== '') {
        await cleanUpUnfinishedIssues(client, [solidityReviewKey])
      }
      throw error
    }
  }

  async function cloneIssue(client: jira.Version3Client, originalIssueKey: string, projectKey: string): Promise<string> {
    try {
      core.debug(`Trying to clone ${originalIssueKey}`)
      const originalIssue = await client.issues.getIssue({ issueIdOrKey: originalIssueKey });

      if (originalIssue.fields.issuetype === undefined) {
        throw new Error(`Issue ${originalIssueKey} is missing issue type id. This should not happen.`)
      }

      const newIssue = await client.issues.createIssue({
        fields: {
          project: {
            key: projectKey,
          },
          priority: originalIssue.fields.priority,
          summary: originalIssue.fields.summary,
          description: originalIssue.fields.description,
          issuetype: { id: originalIssue.fields.issuetype.id },
        },
      });
      core.debug(`Cloned issue key: ${newIssue.key}`)
      return newIssue.key;
    } catch (error) {
      core.error(`Error cloning issue ${originalIssueKey}: ` + error)
      throw error
    }
  }

  async function cloneLinkedIssues(client: jira.Version3Client, projectKey: string, originalIssueKey: string, newIssueKey: string, issueTypes: string[], expectedLinkedIssues?: number) {
    const linkedIssuesKeys: string[] = []
    try {
        core.debug(`Cloning to ${newIssueKey} all issues with type '${join(...issueTypes)}' linked to ${originalIssueKey}`)
      const originalIssue = await client.issues.getIssue({ issueIdOrKey: originalIssueKey });

      // Check the issue's links for any linked issues
      const linkedIssues = originalIssue.fields.issuelinks.filter(link => {
        const issueTypeName = link.inwardIssue?.fields?.issuetype?.name
        if (!issueTypeName) {
          return false
        }
        return issueTypes.length === 0 || issueTypes.includes(issueTypeName)
      });

      if (expectedLinkedIssues && linkedIssues.length !== expectedLinkedIssues) {
        throw new Error(`Expected exactly ${expectedLinkedIssues} linked issues of type ${join(...issueTypes)}, but got ${linkedIssues.length}`)
      }

      for (const issueLink of linkedIssues) {
        if (!issueLink.inwardIssue?.key) {
          throw new Error(`Issue link ${issueLink.id} was missing inward issue or inward issue key`)
        }

        const linkedIssue = await client.issues.getIssue({ issueIdOrKey: issueLink.inwardIssue?.key });

        if (linkedIssue.fields.issuetype === undefined) {
            throw new Error(`Issue ${linkedIssue.key} is missing issue type id. This should not happen.`)
          }

        const newLinkedIssue = await client.issues.createIssue({
          fields: {
            project: {
              key: projectKey,
            },
            priority: linkedIssue.fields.priority,
            summary: linkedIssue.fields.summary,
            description: linkedIssue.fields.description,
            issuetype: { id: linkedIssue.fields.issuetype.id },
          },
        });
        linkedIssuesKeys.push(newLinkedIssue.key)

        core.debug(`Cloned linked issue key: ${newLinkedIssue.key}`);

        copyAllChecklists(linkedIssue.id, newLinkedIssue.id)

        await client.issueLinks.linkIssues({
          type: { name: 'Blocks' },
          inwardIssue: { key: newLinkedIssue.key },
          outwardIssue: { key: newIssueKey },
        });

        core.debug(`Linked ${newLinkedIssue.key} to issue ${newIssueKey}`);
      }
    } catch (error) {
        core.error(`Error cloning linked issues from ${originalIssueKey} to ${newIssueKey}:  ${error}`);
        core.info('issues so far: ' + linkedIssuesKeys)
        await cleanUpUnfinishedIssues(client, linkedIssuesKeys)
        throw error
    }
  }

  async function cleanUpUnfinishedIssues(client: jira.Version3Client, issueKeys: string[]): Promise<unknown> {
    try {
      for (const key of issueKeys) {
        await closeIssue(client, key, 'Closing issue due to an error in automatic creation of Solidity Review')
      }
      return
    } catch (error) {
      core.error(`Failed to close at least one of issues: ${join(...issueKeys)} due to: ${error}. Please close them manually`)
      return error
    }
  }

  async function closeIssue(client: jira.Version3Client, issueKey: string, commentText: string) {
    // in our JIRA '81' is transitionId of `Closed` status, using transition name did not work
    return transitionIssueWithComment(client,issueKey, '81', 'Declined', commentText)
  }

  /**
   * @returns void
   * @throws {Error} If it fails to close the issue.
   */
  async function transitionIssueWithComment(client: jira.Version3Client, issueKey: string, transitionId: string, resolution: string, commentText: string) {
    try {
      await client.issues.doTransition({
        issueIdOrKey: issueKey,
        transition: {
          id: transitionId
        },
        fields: {
          resolution: {
            name: resolution
          }
        },
        update: {
          comment: [
            {
              add: {
                body: {
                  type: 'doc',
                  version: 1,
                  content: [
                    {
                      type: 'paragraph',
                      content: [
                        {
                          type: 'text',
                          text: commentText
                        }
                      ]
                    }
                  ]
                }
              }
            }
          ]
        }
      });

      core.debug(`Issue ${issueKey} successfully closed with comment.`);
    } catch (error) {
      core.error(`Failed to update issue ${issueKey}: ${error}`);
      throw error
    }
  }

  async function copyAllChecklists(sourceIssueId: string, targetIssueId: string) {
    core.debug(`Copying all checklists from ${sourceIssueId} to ${targetIssueId}`)
    const checklistProperty = 'sd-checklists-0'
    const checklistJson = await getChecklistsFromIssue(sourceIssueId, checklistProperty)
    addChecklistsToIssue(targetIssueId, checklistProperty, checklistJson)
    core.debug(`Copied all checklists from ${sourceIssueId} to ${targetIssueId}`)
  }

  async function addChecklistsToIssue(issueId: string, checklistProperty: string, checklistsJson: JSON) {
    core.debug(`Adding checklists to issue ${issueId}`)
    const { jiraHost, jiraUserName, jiraApiToken } = getJiraEnvVars();

    try {
    await axios.put(
        `${jiraHost}rest/api/3/issue/${issueId}/properties/${checklistProperty}`,
        checklistsJson,
        {
        auth: {
            username: jiraUserName,
            password: jiraApiToken,
        }
    },
    );
    core.debug(`Added checklists successfully`)
    } catch (error) {
        core.error(`Failed to add checklists to issue ${issueId}: ${error}`)
        throw error
    }
  }

  async function getChecklistsFromIssue(issueId: string, checklistProperty: string): Promise<JSON> {
    core.debug(`Fetching all checklists from issue ${issueId}`)
    const { jiraHost, jiraUserName, jiraApiToken } = getJiraEnvVars();

    try {
        const response = await axios.get(
          `${jiraHost}rest/api/3/issue/${issueId}/properties/${checklistProperty}`,
          {
            auth: {
              username: jiraUserName,
              password: jiraApiToken,
            },
          }
        );

        if (response.data.value?.checklists && (response.data.value?.checklists as Array<JSON>).length > 0) {
            core.debug(`Found ${(response.data.value?.checklists as Array<JSON>).length} checklists`)
            return response.data.value as JSON
        }

        throw new Error('Checklist response had unexpected content: ' + response.data)

      } catch (error) {
        core.error(`Error reading checklists from issueId ${issueId}: ${error}`);
        throw error
      }
  }

  async function getOpenSolidityReviewIssuesForProject(client: jira.Version3Client, projectKey: string, issueKeysToIgnore: string[]): Promise<string[]> {
    //TODO: change 'Initiative' to 'Solidity Review' once it has been created
    const issueKeys = await getIssueKeys(client, projectKey, 'Initiative', 'Solidity Review', 'Open', issueKeysToIgnore, 10)
    core.info(`Found ${issueKeys.length} open Solidity Review issues for project '${projectKey}'`)
    return issueKeys
  }

  function extracProjectFromIssueKey(issueKey: string): string | undefined {
    const pattern = /([A-Z]{2,})-\d+/

    const match = issueKey.toUpperCase().match(pattern);
    const projectExtracted = match ? match[1] : undefined

    core.debug(`Extracted following project '${projectExtracted}' from issue '${issueKey}'`)

    return projectExtracted
  }

  function readSolidityReviewTemplateKey(): string {
    const issueKey = process.env.SOLIDITY_REVIEW_TEMPLATE_KEY;
    if (!issueKey) {
      throw Error("Missing required environment variable SOLIDITY_REVIEW_TEMPLATE_KEY");
    }

    return issueKey
  }

  async function run() {
    try {
      await main();
    } catch (error) {
      if (error instanceof Error) {
        return core.setFailed(error.message);
      }
      core.setFailed(error as any);
    }
  }

  run();