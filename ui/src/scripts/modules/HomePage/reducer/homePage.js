import useActions from 'scripts/hooks/useActions';

const types = {
    HOME_PAGE_ON_RESET_UI: 'HOME_PAGE_ON_RESET_UI',
    HOME_PAGE_FORM_CHANGE: 'HOME_PAGE_FORM_CHANGE',
    HOME_PAGE_SHORTEN_URL: 'HOME_PAGE_SHORTEN_URL',
};

const initialState = {
    isShortened: false,
    longUrl: '',
    shortUrl: '',
    redirectUrl: '',
};

function reducer(state = initialState, action) {
    switch (action.type) {
        case types.HOME_PAGE_ON_RESET_UI: {
            return initialState;
        }

        case types.HOME_PAGE_FORM_CHANGE: {
            return {
                ...state,
                [action.name]: action.value,
            };
        }

        case types.HOME_PAGE_SHORTEN_URL: {
            return {
                ...state,
                shortUrl: action.shortUrl,
                longUrl: action.redirectUrl,
                isShortened: true,
            };
        }

        default: {
            return state;
        }
    }
}

const actions = {
    onResetUI: () => ({
        type: types.HOME_PAGE_ON_RESET_UI,
    }),

    onFormChange: (name, value) => ({
        type: types.HOME_PAGE_FORM_CHANGE,
        name,
        value,
    }),

    onShortenUrl: (shortUrl, redirectUrl) => ({
        type: types.HOME_PAGE_SHORTEN_URL,
        shortUrl,
        redirectUrl,
    }),
};

const useHomePageActions = () => useActions(actions);

export { useHomePageActions };
export default reducer;
