/**
 * Copyright 2002-2016 Evgeny Gryaznov
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package org.textmapper.lapg.builder;

import org.textmapper.lapg.api.DerivedSourceElement;
import org.textmapper.lapg.api.SourceElement;
import org.textmapper.lapg.api.Symbol;
import org.textmapper.lapg.api.ast.AstType;

abstract class LiSymbol extends LiUserDataHolder implements Symbol, DerivedSourceElement {

	private String name;
	// OR
	private String nameHint;

	private int index;
	protected final SourceElement origin;
	private AstType mapping;
	private boolean unused;

	protected LiSymbol(String name, boolean isAnonymous, SourceElement origin) {
		this.name = isAnonymous ? null : name;
		this.nameHint = isAnonymous ? name : null;
		this.origin = origin;
	}

	@Override
	public int getIndex() {
		return index;
	}

	void setIndex(int index) {
		this.index = index;
	}

	@Override
	public AstType getType() {
		return mapping;
	}

	void setType(AstType mapping) {
		this.mapping = mapping;
	}

	@Override
	public String getName() {
		return name;
	}

	@Override
	public void setName(String value) {
		if (name != null) throw new IllegalStateException();

		name = value;
	}

	@Override
	public void updateNameHint(String nameHint) {
		this.nameHint = nameHint;
	}

	@Override
	public String getNameHint() {
		return nameHint;
	}

	@Override
	public boolean isTerm() {
		return false;
	}

	@Override
	public boolean isUnused() {
		return unused;
	}

	@Override
	public SourceElement getOrigin() {
		return origin;
	}

	@Override
	public String toString() {
		return LiUtil.getSymbolName(this);
	}

	void setUnused() {
		unused = true;
	}
}
