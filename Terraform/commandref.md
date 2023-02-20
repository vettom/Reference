# Commands to remember

#Plan
### -replace
 Plan to replace a specific resource.

#refresh
 Refresh state file with current status of resources in real world. Ideal when someone makes changes outside TF refresh will help TF to update state file
 > Outputs will exist as they are stored separate from resource. Replaced at next apply
> TF perform refresh during plan, apply and destroy


#state 
### mv
 Can be used to move resource from one state file to other or to rename. This updates state file only not the config files
 Can be done without destroying and recreating resources
 Must update config files to reflect changes, else tate file will not match config and could result in creation or destruction of resoure.
 ### rm  
   Remove resource from statfile only
### import
  Import resource to state file
