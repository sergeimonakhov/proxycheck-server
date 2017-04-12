commit:
	git add .
	git commit -m "$m"
	git push -u origin $b

tag:
	git tag "$t"
	git push origin --tags
