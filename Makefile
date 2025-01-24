# Targets
.PHONY: orch core create-branch create-branch-from-main perf-test

# Run the Core service
core:
	@echo "Running Core service"
	cd core && go run main.go

core-local:
	@echo "Running Core service"
	cd core && PORT=8002 go run main.go

# Create a new branch
create-branch:
	@read -p "Enter the new branch name: " branch_name; \
	if [ -z "$$branch_name" ]; then \
		echo "Branch name is required."; \
		exit 1; \
	fi; \
	git checkout -b $$branch_name

# Create a new branch from main
create-branch-from-main:
	@read -p "Enter the new branch name: " branch_name; \
	if [ -z "$$branch_name" ]; then \
		echo "Branch name is required."; \
		exit 1; \
	fi; \
	git checkout main && git pull origin main && git checkout -b $$branch_name