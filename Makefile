commit:
	git add .
	git commit -m "$m"
	git push -u origin $b

tag:
	git tag -d latest
	git push origin :refs/tags/latest
	git fetch --tags
	git tag $tag
	git tag latest
	git push origin --tags
