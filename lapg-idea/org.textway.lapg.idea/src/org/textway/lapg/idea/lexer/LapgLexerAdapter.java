/**
 * Copyright (c) 2010-2011 Evgeny Gryaznov
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see http://www.gnu.org/licenses/.
 */
package org.textway.lapg.idea.lexer;

import com.intellij.lexer.LexerBase;
import com.intellij.psi.TokenType;
import com.intellij.psi.tree.IElementType;
import org.textway.lapg.parser.LapgLexer;
import org.textway.lapg.parser.LapgLexer.LapgSymbol;
import org.textway.lapg.parser.LapgLexer.Lexems;

import java.io.IOException;
import java.io.Reader;
import java.io.StringReader;

public class LapgLexerAdapter extends LexerBase implements LapgTokenTypes {

	private CharSequence myText;
	private LapgLexer lexer;
	private LapgSymbol lexem;
	private int fDocumentLength;
	private int fTokenOffset;
	private int fTokenLength;
	private IElementType current;

	public LapgLexerAdapter() {
	}

	public void start(final CharSequence buffer, int startOffset, int endOffset, int initialState) {
		myText = buffer;
		fDocumentLength = endOffset;
		Reader reader = new StringReader(buffer.toString().substring(startOffset, endOffset));

		try {
			if (lexer == null) {
				lexer = new IdeaLapgLexer(reader);
			} else {
				lexer.reset(reader);
			}
		} catch (IOException ex) {
			/* never happens */
		}
		lexer.setOffset(startOffset);
		fTokenOffset = startOffset;
		fTokenLength = 0;
		lexem = null;
		current = null;
	}

	public int getState() {
		return 0;
	}

	public IElementType getTokenType() {
		locateToken();
		return current;
	}

	public int getTokenStart() {
		locateToken();
		return fTokenOffset;
	}

	public int getTokenEnd() {
		locateToken();
		return fTokenOffset + fTokenLength;
	}

	public void advance() {
		locateToken();
		current = null;
	}

	public CharSequence getBufferSequence() {
		return myText;
	}

	public int getBufferEnd() {
		return fDocumentLength;
	}

	private void locateToken() {
		if (current == null) {
			current = nextToken();
		}
	}

	public IElementType nextToken() {
		fTokenOffset += fTokenLength;
		if (lexem == null) {
			readNext();
		}
		if (fTokenOffset < lexem.offset) {
			fTokenLength = lexem.offset - fTokenOffset;
			return TokenType.BAD_CHARACTER;
		}
		int token = lexem.lexem;
		fTokenLength = lexem.endoffset - fTokenOffset;
		LapgSymbol currentLexem = lexem;
		lexem = null;
		switch (token) {
			case Lexems.code:
				return TOKEN_ACTION;
			case Lexems._skip:
				return WHITESPACE;
			case Lexems._skip_comment:
				return COMMENT;
			case Lexems.scon:
				return STRING;
			case Lexems.icon:
				return ICON;
			case Lexems.identifier:
				return IDENTIFIER;
			case Lexems.regexp:
				return REGEXP;

			// operators
			case Lexems.PERCENT:
				return OP_PERCENT;
			case Lexems.COLONCOLONEQUAL:
				return OP_CCEQ;
			case Lexems.OR:
				return OP_OR;
			case Lexems.EQUAL:
				return OP_EQ;
			case Lexems.EQUALGREATER:
				return OP_EQGT;
			case Lexems.SEMICOLON:
				return OP_SEMICOLON;
			case Lexems.DOT:
				return OP_DOT;
			case Lexems.COMMA:
				return OP_COMMA;
			case Lexems.COLON:
				return OP_COLON;
			case Lexems.LSQUARE:
				return OP_LBRACKET;
			case Lexems.RSQUARE:
				return OP_RBRACKET;
			case Lexems.LPAREN:
				return OP_LPAREN;
			case Lexems.RPAREN:
				return OP_RPAREN;
			case Lexems.LESS:
				return OP_LT;
			case Lexems.GREATER:
				return OP_GT;
			case Lexems.MULT:
				return OP_STAR;
			case Lexems.PLUS:
				return OP_PLUS;
			case Lexems.QUESTIONMARK:
				return OP_QMARK;
			case Lexems.AMPERSAND:
				return OP_AND;
			case Lexems.ATSIGN:
				return OP_AT;

			// keywords
			case Lexems.Ltrue:
				return KW_TRUE;
			case Lexems.Lfalse:
				return KW_FALSE;
			case Lexems.Lprio:
				return KW_PRIO;
			case Lexems.Lshift:
				return KW_SHIFT;
			case Lexems.Lreduce:
				return KW_REDUCE;
			case Lexems.Linput:
				return KW_INPUT;
			case Lexems.Lleft:
				return KW_LEFT;
			case Lexems.Lright:
				return KW_RIGHT;
			case Lexems.Lnonassoc:
				return KW_NONASSOC;
			case Lexems.Lnoeoi:
				return KW_NOEOI;
			case Lexems.Lsoft:
				return KW_SOFT;
			case Lexems.Lclass:
				return KW_CLASS;
			case Lexems.Lspace:
				return KW_SPACE;
		}

		/* default, eoi */
		lexem = currentLexem;
		assert lexem.lexem == Lexems.eoi;
		if (lexem.endoffset < fDocumentLength) {
			fTokenLength = fDocumentLength - fTokenOffset;
			lexem.offset = lexem.endoffset = fDocumentLength;
			return TEMPLATES;
		}
		return null;
	}

	private void readNext() {
		try {
			lexem = lexer.next();
		} catch (IOException e) {
			/* never happens */
		}
	}

	private static class IdeaLapgLexer extends LapgLexer {
		public IdeaLapgLexer(Reader stream) throws IOException {
			super(stream, new ErrorReporter() {
				public void error(int start, int end, int line, String s) {
				}
			});
		}

		@Override
		protected boolean createToken(LapgSymbol lapg_n, int lexemIndex) throws IOException {
			super.createToken(lapg_n, lexemIndex);
			return true;
		}
	}
}
