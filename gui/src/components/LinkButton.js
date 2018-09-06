import React from 'react'
import StyledButton from '@material-ui/core/Button'
import classNames from 'classnames'
import { withStyles } from '@material-ui/core/styles'

const styles = theme => ({
  linkButton: {
    color: theme.palette.common.white,
    backgroundColor: 'transparent',
    textDecoration: 'underline',
    textTransform: 'capitalize',
    fontSize: 'inherit',
    lineHeight: 'inherit',
    '&:hover': {
      backgroundColor: 'transparent',
      textDecoration: 'underline'
    }
  }
})

const LinkButton = ({classes, className, ...props}) => (
  <StyledButton className={classNames(classes.linkButton, className)} {...props} />
)

export default withStyles(styles)(LinkButton)
