package com.mageddo.db;

/**
 * Created by elvis on 29/04/17.
 */
public class DefaultTransactionDefinition  extends org.springframework.transaction.support.DefaultTransactionDefinition {

	public DefaultTransactionDefinition() {
		this(PROPAGATION_REQUIRED, ISOLATION_READ_COMMITTED);
	}
	public DefaultTransactionDefinition(final int propagation, final int isolation) {
		setPropagationBehavior(propagation);
		setIsolationLevel(isolation);
	}
}
