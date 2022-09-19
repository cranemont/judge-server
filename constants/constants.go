package constants

const (
	Success = 1 + iota
	Fail
)

const DIR_NAME_LEN = 16
const MAX_SUBMISSION = 10
const EVENT_CHAN_SIZE = 10
const TASK_EXEC = "Execute"
const TASK_EXITED = "Exited"
const PUBLISH_RESULT = "Publish"

// const BASE_DIR = "./results"
const SANDBOX_BASE_DEV = "/go/src/github.com/cranemont/judge-manager"
const RESULT_PATH_DEV = SANDBOX_BASE_DEV + "/results"
const LIBJUDGER_PATH_DEV = "/app/sandbox/libjudger.so"
const JAVA_POLICY_PATH_DEV = SANDBOX_BASE_DEV + "/policy/java_policy"

// const BASE_DIR = "/go/src/workspace/results"
const BASE_FILE_MODE = 0711

const SANDBOX_BASE_PROD = "/app/sandbox"
const RESULT_PATH_PROD = SANDBOX_BASE_PROD + "/results"
const LIBJUDGER_PATH_PROD = SANDBOX_BASE_PROD + "/libjudger.so"
const JAVA_POLICY_PATH_PROD = SANDBOX_BASE_PROD + "/policy/java_policy"

// FIXME: logger 구현 후 다시 설정
const (
	SANDBOX_LOG_BASE = "/app/sandbox/logs"
	COMPILE_LOG_PATH = SANDBOX_LOG_BASE + "/compile/log.out"
	RUN_LOG_PATH     = SANDBOX_LOG_BASE + "/run/log.out"
	COMPILE_OUT_FILE = "compile.out"
)

const (
	LOG_PATH_DEV  = "./logs/server.log"
	LOG_PATH_PROD = "./logs/server.log" // "/app/logs/server.log"
)

const MAX_MQ_CHANNEL = 10

const TESTCASE_GET_TIMEOUT = 10
const TOKEN_HEADER = "judge-server-token"

const (
	EXCHANGE = "judger-exchange"
	RESULT_KEY = "result"
)