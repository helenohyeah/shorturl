import React from 'react';
import { connect } from 'react-redux';

// const selector = state => state.app;
const mapStateToProps = state => ({
    //     isPasswordHidden: selector(state).isPasswordHidden,
    //     username: selector(state).username,
    //     password: selector(state).password,
});

const UserPageContainer = ({}) => {
    return <div>User page</div>;
};

// export default connect(mapStateToProps)(UserPageContainer);
export default UserPageContainer;
