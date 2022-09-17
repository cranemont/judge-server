package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/cranemont/judge-manager/file"
	"github.com/cranemont/judge-manager/handler/judge"
	"github.com/cranemont/judge-manager/logger"
	"github.com/cranemont/judge-manager/sandbox"
)

var handler = "JudgeHandler"

type JudgeResposne struct {
	StatusCode Code                  `json:"statusCode"` // handler's status code
	Data       judge.JudgeTaskResult `json:"data"`
}

type JudgeRequest struct {
	Code        string `json:"code"`
	Language    string `json:"language"`
	ProblemId   int    `json:"problemId"`
	TimeLimit   int    `json:"timeLimit"`
	MemoryLimit int    `json:"memoryLimit"`
}

type JudgeHandler struct {
	langConfig sandbox.LangConfig
	file       file.FileManager
	judger     *judge.Judger
	logging    *logger.Logger
}

func NewJudgeHandler(
	langConfig sandbox.LangConfig,
	file file.FileManager,
	judger *judge.Judger,
	logging *logger.Logger,
) *JudgeHandler {
	return &JudgeHandler{langConfig, file, judger, logging}
}

// handle top layer logical flow
func (j *JudgeHandler) Handle(req JudgeRequest) (result JudgeResposne, err error) {
	res := JudgeResposne{StatusCode: INTERNAL_SERVER_ERROR, Data: judge.JudgeTaskResult{}}
	task := judge.NewTask(
		req.Code, req.Language, strconv.Itoa(req.ProblemId), req.TimeLimit, req.MemoryLimit,
	)
	dir := task.GetDir()

	defer func() {
		j.file.RemoveDir(dir)
		j.logging.Debug(fmt.Sprintf("task %s done: total time: %s", dir, time.Since(task.StartedAt)))
	}()

	if err := j.file.CreateDir(dir); err != nil {
		return res, fmt.Errorf("handler: %s: failed to create base directory: %w", handler, err)
	}

	srcPath, err := j.langConfig.MakeSrcPath(dir, task.GetLanguage())
	if err != nil {
		return res, fmt.Errorf("handler: %s: failed to create src path: %w", handler, err)
	}

	if err := j.file.CreateFile(srcPath, task.GetCode()); err != nil {
		return res, fmt.Errorf("handler: %s: failed to create src file: %w", handler, err)
	}

	err = j.judger.Judge(task)
	if err != nil {
		if errors.Is(err, judge.ErrTestcaseGet) {
			res.StatusCode = TESTCASE_GET_FAILED
		} else if errors.Is(err, judge.ErrCompile) {
			res.StatusCode = COMPILE_ERROR
			res.Data = task.Result
			return res, nil
		} else {
			res.StatusCode = INTERNAL_SERVER_ERROR
		} // run, grade 등 추가
		return res, fmt.Errorf("handler: judge failed: %w", err)
	} else {
		res.StatusCode = SUCCESS
	}

	res.Data = task.Result
	return res, nil
}

func (h *JudgeHandler) ResultToJson(result JudgeResposne) []byte {
	res, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}
	return res
}
