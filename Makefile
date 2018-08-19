.PHONY: doc

doc:
	@echo making doc
	java -jar plantuml.jar \
		doc/cache_uml.txt \
		doc/memory_uml.txt \
		doc/decorator_uml.txt \
		doc/cache.v1_uml.txt
	cd doc; pandoc -f markdown-implicit_figures requirements.md \
		--latex-engine xelatex -o requirements.pdf
	cd doc; pandoc -f markdown-implicit_figures design.md \
		--latex-engine xelatex -o design.pdf
