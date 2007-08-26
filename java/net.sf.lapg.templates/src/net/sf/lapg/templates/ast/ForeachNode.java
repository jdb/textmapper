package net.sf.lapg.templates.ast;

import java.util.Collection;
import java.util.List;

import net.sf.lapg.templates.api.EvaluationException;
import net.sf.lapg.templates.api.IEvaluationEnvironment;

public class ForeachNode extends CompoundNode {
	
	private static final String INDEX = "index";
	String var;
	ExpressionNode selectExpr;
	
	public ForeachNode(String var, ExpressionNode selectExpr) {
		this.var = var;
		this.selectExpr = selectExpr;
	}

	protected void emit(StringBuffer sb, Object context, IEvaluationEnvironment env) {
		try {
			Object select = env.evaluate(selectExpr, context, false);
			Object prevVar = env.getVariable(var);
			Object prevIndex = env.getVariable(INDEX);
			int index = 0;
			try {
				if( select instanceof Collection ) {
					for( Object o : (Collection)select ) {
						env.setVariable(var, o);
						env.setVariable(INDEX, index++);
						for( Node n : instructions ) {
							n.emit(sb, context, env);
						}
					}
				} else if(select instanceof Object[]) {
					for( Object o : (Object[])select ) {
						env.setVariable(var, o);
						env.setVariable(INDEX, index++);
						for( Node n : instructions ) {
							n.emit(sb, context, env);
						}
					}
				} else {
					env.fireError("In foreach `"+selectExpr.toString()+"` should be Object[] or Collection for " + context.getClass().getCanonicalName());
				}
			} finally {
				env.setVariable(var, prevVar);
				env.setVariable(INDEX, prevIndex);
			}
		} catch( EvaluationException ex ) {
			/* ignore, skip foreach */
		}
	}
}
