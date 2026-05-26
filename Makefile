.PHONY: release release-minor release-major

release:
	@./release.sh patch

release-minor:
	@./release.sh minor

release-major:
	@./release.sh major
