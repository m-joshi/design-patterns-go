GOCMD=go
    GORUN=$(GOCMD) run
    GOCLEAN=$(GOCMD) clean

run_strategy:
	$(GORUN) strategyPattern/main/main.go

run_observer:
	$(GORUN) observerPattern/main/main.go

run_decorator:
	$(GORUN) decoratorPattern/main/main.go

run_scatter_gather:
	$(GORUN) scatterGatherPattern/main/main.go

clean:
	$(GOCLEAN)
