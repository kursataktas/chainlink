import React from 'react'
import Card from '@material-ui/core/Card'
import CardContent from '@material-ui/core/CardContent'

export default ({children, classes, className}) => (
  <Card>
    <CardContent className={className}>
      {children}
    </CardContent>
  </Card>
)
