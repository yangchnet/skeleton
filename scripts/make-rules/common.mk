SHELL := /bin/bash

# project name
ifeq ($(origin PROJECT_NAME),undefined)
	PROJECT_NAME := skeleton
endif

# common password
ifeq ($(origin PASSWORD),undefined)
	PASSWORD := ?@-12
endif

# db type
ifeq ($(origin DB_TYPE),undefined)
	DB_TYPE := mysql
endif

ifeq ($(origin DB_PORT),undefined)
	DB_PORT := 13306
endif


COMMON_SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))

ifeq ($(origin ROOT_DIR),undefined)
ROOT_DIR := $(abspath $(shell cd $(COMMON_SELF_DIR)/../.. && pwd -P))
endif

ifeq ($(origin OUTPUT_DIR),undefined)
OUTPUT_DIR := $(ROOT_DIR)/_output
$(shell mkdir -p $(OUTPUT_DIR))
endif

ifeq ($(origin TOOLS_DIR),undefined)
TOOLS_DIR := $(ROOT_DIR)/tools
$(shell mkdir -p $(TOOLS_DIR))
endif

ifeq ($(origin TMP_DIR),undefined)
TMP_DIR := $(ROOT_DIR)/tmp
$(shell mkdir -p $(TMP_DIR))
endif

# set the version number
ifeq ($(origin VERSION),undefined)
VERSION := $(shell git describe --tags --always --match="v*")
endif

#Check if the tree is dirty, default is dirty
GIT_TREE_STATE :="dirty"
ifeq (,$(shell git status --porcelain 2>/dev/null))
	GIT_TREE_STATE="clean"
endif
GIT_COMMIT := $(shell git rev-parse HEAD)

# Minimum test coverage
ifeq ($(origin COVERAGE),undefined)
COVERAGE := 60
endif

# The OS must be linux when building docker images
PLATFORMS ?= linux_amd64, linux_arm64

# Set a specific PALTFORM
ifeq ($(origin PLATFORM),undefined)
	ifeq ($(origin GOOS),undefined)
		GOOS := $(shell go env GOOS)
	endif
	ifeq ($(origin GOARCH),undefined)
		GOARCH := $(shell go env GOARCH)
	endif
	PLATFORM  := $(GOOS)_$(GOARCH)
	# Use linux as the default OS when building images
	IMAGE_PLAT := linux_$(GOARCH)
else
	GOOS := $(word 1, $(subst _, ,$(PLATFORM)))
	GOARCH := $(word 2, $(subst _, ,$(PLATFORM)))
	IMAGE_PLAT := $(PLATFORM)
endif

# Set GOPATH
ifeq ($(origin GOPATH),undefined)
	GOPATH := $(shell go env GOPATH)
endif

# Set GOBIN
ifeq ($(origin GOBIN),undefined)
	GOBIN := $(addprefix $(GOPATH),/bin)
endif

ifndef V
MAKEFLAGS += --no-print-directory
endif

# Specify tools severity, include: BLOCKER_TOOLS, CRITICAL_TOOLS, TRIVIAL_TOOLS.
# Missing BLOCKER_TOOLS can cause the CI flow execution failed, i.e. `make all` failed.
# Missing CRITICAL_TOOLS can lead to some necessary operations failed. i.e. `make release` failed.
# TRIVIAL_TOOLS are Optional tools, missing these tool have no affect.
BLOCKER_TOOLS ?= goimports golangci-lint
CRITICAL_TOOLS ?= mockgen protoc-gen-go protoc-gen-grpc-gateway \
 protoc-gen-go-grpc protoc-gen-openapiv2 wire buf \
 ent 
TRIVIAL_TOOLS ?= 

COMMA := ,
EMPTY :=
SPACE := $(EMPTY) $(EMPTY) 