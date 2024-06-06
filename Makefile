REGEX_SEMVER=[^0-9.]*\([0-9.]*\).*/\1/
STDV=standard-version
STDV_PREVIEW=${STDV} --dry-run --release-as
STDV_RELEASE=${STDV} --release-as
PREVIEW=standard-version --dry-run --release-as
GET_VERSION=sed -n '/release/ s/${REGEX_SEMVER}p'
GIT_BRANCH=$$(git branch --show-current)
VERSION_FILE=version.go
VERSION_PACKAGE=lib
VERSION=v

push:
	git push -u origin main

preview-%:
	${STDV_PREVIEW} $*

# git push --follow-tags origin heads/v0.2.0
release-%: VERSION=$$(echo "v$$(${STDV_PREVIEW} $* | ${GET_VERSION})")
release-%:
	git checkout -b "release/$(VERSION)"
	@echo "// WARNING: auto-generated by Makefile release target: DO NOT EDIT." > $(VERSION_FILE)
	@echo "" >> $(VERSION_FILE)
	@echo "package $(VERSION_PACKAGE)" >> $(VERSION_FILE)
	@echo "" >> $(VERSION_FILE)
	@echo "const Version=\"$(VERSION)\"" >> $(VERSION_FILE)
	git commit -am "release: $(VERSION)"
	git push -u origin ${GIT_BRANCH}
	${STDV_RELEASE} --release-as $*
	#git push --follow-tags origin $(GIT_BRANCH)