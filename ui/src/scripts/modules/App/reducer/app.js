import useActions from 'scripts/hooks/useActions';

const types = {
    APP_LOGIN: 'APP_LOGIN',
    APP_LOGOUT: 'APP_LOGOUT',
};

const initialState = {
    user: {},
};

function reducer(state = initialState, action) {
    switch (action.type) {
        case types.APP_LOGIN: {
            return {
                ...state,
                user: action.user,
            };
        }

        case types.APP_LOGOUT: {
            return {
                ...state,
                user: {},
            };
        }

        default: {
            return state;
        }
    }
}

const actions = {
    onLogin: user => ({
        type: types.APP_LOGIN,
        user,
    }),

    onLogout: () => ({
        type: types.APP_LOGOUT,
    }),
};

const useAppActions = () => useActions(actions);

export { useAppActions };
export default reducer;
