# Env variables

### Enable dedug/trace
    TF_LOG=DEBUG  : option [Trace, Debug, off, info
    TF_LOG_PATH=./terraform.log  : Path for TF_LOG_PATH. If not set only output to console. 
    TF_INPUT=0  Disable promptes for input equal to -input=false
    TF_VAR_name  used to set environment variables. export TF_VAR_region=us-west-1
    TF_CLI_ARGS : Additional argument that ca =n be used to pass for cli
    TF_CLI_ARGS_name : Arguments for specific names action. TF_CLI_ARGS_plan="-refresh=false"
    TF_DATA_DIR  : to set where terraform keep preworking config, if not .terraform
    TF_WORKSPACE : sets workspace, no need to set
    TF_IN_AUTOMATION : If set TF know it is automation and adjusts output data. no suggestions.
    TF_LOG_PROVIDER  : Enable logging for provider
    TF_LOG_CORE     : Core app logging diagnostic
 


> TF_LOG defaults to trace if set with invalid value. TF_LOG_PATH requires TF_LOG