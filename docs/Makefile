.PHONY: clean

spec.pdf: spec.tex
	latexmk --pdf $<

clean:
	$(RM) *.aux
	$(RM) *.bbl
	$(RM) *.bcf
	$(RM) *.blg
	$(RM) *.fdb_latexmk
	$(RM) *.fls
	$(RM) *.log
	$(RM) *.out
	$(RM) *.pdf
	$(RM) *.run.xml
	$(RM) *.synctex.gz
