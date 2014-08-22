package org.textmapper.tool.importer;

import java.io.*;
import org.textmapper.tool.importer.BisonLexer.ErrorReporter;
import org.textmapper.tool.importer.BisonParser.ParseException;

public class BisonMain {

	private static final String EXTENSION = ".y";
	private static final String ENCODING = "utf-8";

	private int counter = 0;
	private int errors = 0;

	public static String getFileContents(InputStream stream) throws IOException {
		StringBuilder contents = new StringBuilder();
		char[] buffer = new char[2048];
		Reader in = new InputStreamReader(stream, ENCODING);
		try {
			int count;
			while ((count = in.read(buffer)) > 0) {
				contents.append(buffer, 0, count);
			}
		} finally {
			in.close();
		}
		return contents.toString();
	}

	private void parse(File file) {
		System.out.println("parsing " + file.getPath());
		counter++;

		final int[] problems = new int[]{0};
		ErrorReporter reporter = new ErrorReporter() {
			@Override
			public void error(String message, int line, int offset, int endoffset) {
				System.out.println("   " + line + ": " + message);
				problems[0]++;
			}
		};

		try {
			String contents = getFileContents(new FileInputStream(file));
			BisonLexer lexer = new BisonLexer(contents.toCharArray(), reporter);
			BisonParser parser = new BisonParser(reporter);
			parser.parse(lexer);
		} catch (ParseException ex) {
			/* not parsed */
		} catch (IOException ex) {
			System.out.println("   I/O problem: " + ex.getMessage());
			problems[0]++;
		}
		if (problems[0] > 0) {
			System.out.println("not parsed " + file.getPath() + ", " + problems[0] + " problem(s)");
			errors++;
		}
	}

	private void walk(File dir) {
		File[] list = dir.listFiles();
		if (list == null) return;

		for (File f : list) {
			if (f.getName().endsWith(EXTENSION) && f.isFile()) {
				parse(f);
			}

			if (f.isDirectory()) {
				walk(f);
			}
		}
	}

	private void printStatistics() {
		System.out.println("Processed " + counter + " files with " + errors + " errors.");
	}

	public static void main(String[] args) {
		if (args.length == 0) {
			System.out.println("Provide a list of directories with " + EXTENSION + " files.");
			return;
		}

		BisonMain instance = new BisonMain();
		for (String path : args) {
			instance.walk(new File(path));
		}
		instance.printStatistics();
	}
}
