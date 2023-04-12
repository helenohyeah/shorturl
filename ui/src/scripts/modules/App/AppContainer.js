import React from 'react';
import { connect } from 'react-redux';
import Header from 'scripts/components/Header';

const selector = state => state.app;
const mapStateToProps = state => ({
    // isPasswordHidden: selector(state).isPasswordHidden,
    // username: selector(state).username,
    // password: selector(state).password,
});

const AppContainer = ({
    // isPasswordHidden,
    // username,
    // password,
    children,
}) => {
    return (
        <div>
            <Header />
            <div>{children}</div>
        </div>
    );
};

export default connect(mapStateToProps)(AppContainer);
