GOCMD=go
    GOBUILD=$(GOCMD) build
    GORUN=$(GOCMD) run
    GOCLEAN=$(GOCMD) clean

    BINARY_NAME= "decorator_pattern"

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

run_strategy:
	./strategy_pattern

run_observer:
	./observer_pattern

run_decorator:
	./decorator_pattern

run_scatter_gather:
	$(GORUN) scatterGatherPattern/main/main.go

clean:
	$(GOCLEAN)
