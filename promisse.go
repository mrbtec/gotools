package promise

type Result interface{}
type Promise struct {
    resultChan chan Result
    errChan    chan error
}

func New(fn func() (Result, error)) *Promise {
    p := &Promise{
        resultChan: make(chan Result),
        errChan:    make(chan error),
    }

    go func() {
        res, err := fn()
        if err != nil {
            p.errChan <- err
            return
        }
        p.resultChan <- res
    }()

    return p
}

func (p *Promise) Then(fn func(Result) (Result, error)) *Promise {
    newPromise := &Promise{
        resultChan: make(chan Result),
        errChan:    make(chan error),
    }

    go func() {
        res := <-p.resultChan
        newRes, err := fn(res)
        if err != nil {
            newPromise.errChan <- err
            return
        }
        newPromise.resultChan <- newRes
    }()

    return newPromise
}

func (p *Promise) Catch(fn func(error) error) *Promise {
    newPromise := &Promise{
        resultChan: make(chan Result),
        errChan:    make(chan error),
    }

    go func() {
        err := <-p.errChan
        newErr := fn(err)
        newPromise.errChan <- newErr
    }()

    return newPromise
}

func (p *Promise) Wait() {
    select {
    case <-p.resultChan:
    case <-p.errChan:
    }
}

func (p *Promise) Result() (Result, error) {
    select {
    case res := <-p.resultChan:
        return res, nil
    case err := <-p.errChan:
        return nil, err
    }
}
