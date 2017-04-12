commit:
	git add .
	git commit -m "$m"
	git push -u origin $b

tag:
	git tag $tag
	git push origin --tags
