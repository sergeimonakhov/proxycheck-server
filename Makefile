commit:
	git add .
	git commit -m "$m"
	git push -u origin "$branch"

tag:
	git tag -a "$a" -m "$m"
	git push origin --tags
