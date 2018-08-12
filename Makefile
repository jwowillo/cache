.PHONY: doc

doc:
	@echo making doc
	cd doc; pandoc requirements.md --latex-engine xelatex \
		-o requirements.pdf
