from PyPDF2 import PdfFileWriter, PdfFileReader, PdfFileMerger
import fire
import glob


def split(filename):
    inputpdf = PdfFileReader(open(filename, "rb"))

    for i in range(inputpdf.numPages):
        output = PdfFileWriter()
        output.addPage(inputpdf.getPage(i))
        with open("document-page-%s.pdf" % i, "wb") as outputStream:
            output.write(outputStream)
    print("done")


def merge_pdf(path="", output=""):
    
    filenames = glob.glob(path + "*.pdf")
    
    merger = PdfFileMerger()
    for filename in filenames:
        merger.append(PdfFileReader(open(filename, 'rb')))

    merger.write(output)

if __name__ == "__main__":
    fire.Fire()
