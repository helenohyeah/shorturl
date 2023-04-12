import React from 'react';
import { connect } from 'react-redux';

// const selector = state => state.app;
const mapStateToProps = state => ({
    //     isPasswordHidden: selector(state).isPasswordHidden,
    //     username: selector(state).username,
    //     password: selector(state).password,
});

const NotFoundContainer = ({}) => {
    return <div>404 not found</div>;
};

// export default connect(mapStateToProps)(NotFoundContainer);
export default NotFoundContainer;
