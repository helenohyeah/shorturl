import useActions from 'scripts/hooks/useActions';
import { of } from 'await-of';
import * as API from 'request';

const types = {
    HOME_PAGE_ON_RESET_UI: 'HOME_PAGE_ON_RESET_UI',
    HOME_PAGE_FORM_CHANGE: 'HOME_PAGE_FORM_CHANGE',
    HOME_PAGE_SHORTEN_URL_SUCCESS: 'HOME_PAGE_SHORTEN_URL_SUCCESS',
    HOME_PAGE_SHORTEN_URL_ERROR: 'HOME_PAGE_SHORTEN_URL_ERROR',
};

const initialState = {
    isShortened: false,
    longUrl: '',
    shortUrl: '',
    redirectUrl: '',
    error: '',
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

        case types.HOME_PAGE_SHORTEN_URL_ERROR: {
            return {
                ...state,
                error: action.error,
            };
        }

        case types.HOME_PAGE_SHORTEN_URL_SUCCESS: {
            return {
                ...state,
                shortUrl: action.encodedUrl,
                longUrl: action.redirectUrl,
                isShortened: true,
                error: '',
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

    onShortenUrl: (userId, longUrl) => async dispatch => {
        const [resp = {}, error] = await of(
            API.shortenUrl({
                userId,
                originalUrl: longUrl,
            })
        );
        if (error) {
            dispatch({
                type: types.HOME_PAGE_SHORTEN_URL_ERROR,
                error: error.message,
            });
            throw error;
        }
        const { data: url = {} } = resp;
        dispatch({
            type: types.HOME_PAGE_SHORTEN_URL_SUCCESS,
            encodedUrl: url.encodedUrl,
            redirectUrl: url.redirectUrl,
        });
    },
};

const useHomePageActions = () => useActions(actions);

export { useHomePageActions };
export default reducer;
