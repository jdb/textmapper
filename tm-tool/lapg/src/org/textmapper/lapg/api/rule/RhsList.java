package org.textmapper.lapg.api.rule;

/**
 * evgeny, 1/3/13
 */
public interface RhsList extends RhsRoot {

	RhsPart getElement();

	RhsPart getCustomInitialElement();

	RhsPart getSeparator();

	boolean isNonEmpty();

	boolean isRightRecursive();
}
