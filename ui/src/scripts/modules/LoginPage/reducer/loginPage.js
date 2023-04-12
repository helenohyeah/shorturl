import useActions from 'scripts/hooks/useActions';

const types = {
    LOGIN_PAGE_RESET_UI: 'LOGIN_PAGE_RESET_UI',
    LOGIN_PAGE_FORM_CHANGE: 'LOGIN_PAGE_FORM_CHANGE',
    LOGIN_PAGE_TOGGLE_SHOW_PASSWORD: 'LOGIN_PAGE_TOGGLE_SHOW_PASSWORD',
};

const initialState = {
    isPasswordHidden: true,
    username: '',
    password: '',
};

function reducer(state = initialState, action) {
    switch (action.type) {
        case types.LOGIN_PAGE_FORM_CHANGE: {
            return {
                ...state,
                [action.name]: action.value,
            };
        }

        case types.LOGIN_PAGE_TOGGLE_SHOW_PASSWORD: {
            return {
                ...state,
                isPasswordHidden: !state.isPasswordHidden,
            };
        }

        case types.LOGIN_PAGE_RESET_UI:
            return initialState;

        default: {
            return state;
        }
    }
}

const actions = {
    onResetUI: () => ({
        type: types.LOGIN_PAGE_RESET_UI,
    }),

    onFormChange: (name, value) => ({
        type: types.LOGIN_PAGE_FORM_CHANGE,
        name,
        value,
    }),

    onToggleShowPassword: () => ({
        type: types.LOGIN_PAGE_TOGGLE_SHOW_PASSWORD,
    }),
};

const useLoginPageActions = () => useActions(actions);

export { useLoginPageActions };
export default reducer;
