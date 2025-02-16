# 获取当前git仓库的commit id
GITCOMMIT := $(shell git rev-parse HEAD)
# 获取当前git提交的时间戳
GITDATE := $(shell git show -s --format='%ct')

LDFLAGSSTRING +=-X main.GitCommit=$(GITCOMMIT)
LDFLAGSSTRING +=-X main.GitDate=$(GITDATE)
LDFLAGSSTRING +=-X main.GitVersion=$(GITVERSION)
LDFLAGS :=-ldflags "$(LDFLAGSSTRING)"

VRF_ABI_ARTIFACT := ./abis/VerifyRandVRF.sol/VerifyRandVRF.json
FACTORY_ABI_ARTIFACT := ./abis/VerifyRandVRFFactory.sol/VerifyRandVRFFactory.json

VRFService:
	env GO111MODULE=on go build $(LDFLAGS) ./cmd/VRFService

clean:
	rm VRFService

test:
	go test -v ./...

lint:
	golangci-lint run ./...

fmt::
	go fmt ./...

tidy::
	go mod tidy -v

# 通过 abigen 工具生成 Go 语言绑定文件，用于与以太坊智能合约进行交互。
#bindings: binding-vrf binding-factory
#
#
#binding-vrf:
#	$(eval temp := $(shell mktemp))
#
#	cat $(VRF_ABI_ARTIFACT) \
#    	| jq -r .bytecode.object > $(temp)
#
#	cat $(VRF_ABI_ARTIFACT) \
#		| jq .abi \
#		| abigen --pkg bindings \
#		--abi - \
#		--out bindings/verifyrandvrf.go \
#		--type VerifyRandVRF \
#		--bin $(temp)
#
#		rm $(temp)
#
#binding-factory:
#	$(eval temp := $(shell mktemp))
#
#	cat $(FACTORY_ABI_ARTIFACT) \
#    	| jq -r .bytecode.object > $(temp)
#
#	cat $(FACTORY_ABI_ARTIFACT) \
#		| jq .abi \
#		| abigen --pkg bindings \
#		--abi - \
#		--out bindings/verifyrandvrffactory.go \
#		--type VerifyRandVRFFactory \
#		--bin $(temp)
#
#		rm $(temp)

# 在 Makefile 中，.PHONY 用于声明某些目标是伪目标，意思是这些目标不会与文件名冲突，并且不依赖于文件系统中的实际文件。这样，无论是否存在同名文件，make 都会执行相应的规则。
.PHONY: \
	VRFService \
	bindings \
	binding-vrf \
	binding-factory \
	clean \
	test \
	lint


# 定义临时文件和常用变量
TEMP_FILE := $(shell mktemp)

# 定义生成 Go 绑定的规则
define generate_binding
    cat $1 | jq -r .bytecode.object > $(TEMP_FILE)
    cat $1 | jq .abi | abigen --pkg bindings --abi - --out $(2) --type $(3) --bin $(TEMP_FILE)
    rm $(TEMP_FILE)
endef

# 绑定 VRF 合约
binding-vrf:
	$(call generate_binding, $(VRF_ABI_ARTIFACT), bindings/verifyrandvrf.go, VerifyRandVRF)

# 绑定 Factory 合约
binding-factory:
	$(call generate_binding, $(FACTORY_ABI_ARTIFACT), bindings/verifyrandvrffactory.go, VerifyRandVRFFactory)

# 主要目标 https://geth.ethereum.org/docs/tools/abigen
bindings: binding-vrf binding-factory

.PHONY: \
	VRFService \
	bindings \
	binding-vrf \
	binding-factory \
	clean \
	test \
	lint
