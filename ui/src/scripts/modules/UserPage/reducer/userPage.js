import useActions from 'scripts/hooks/useActions';
import { of } from 'await-of';
import * as API from 'request';

const types = {
    USER_PAGE_RESET_UI: 'USER_PAGE_RESET_UI',
    USER_PAGE_GET_URLS_SUCCESS: 'USER_PAGE_GET_URLS_SUCCESS',
    USER_PAGE_GET_URLS_ERROR: 'USER_PAGE_GET_URLS_ERROR',
};

const initialState = {
    urls: [],
};

function reducer(state = initialState, action) {
    switch (action.type) {
        case types.USER_PAGE_RESET_UI: {
            return initialState;
        }

        case types.USER_PAGE_GET_URLS_SUCCESS: {
            return { ...state, urls: action.urls };
        }

        default: {
            return state;
        }
    }
}

const actions = {
    onResetUI: () => ({
        type: types.USER_PAGE_RESET_UI,
    }),

    onGetUrls: userId => async dispatch => {
        const [resp = {}, error] = await of(API.getUrlsByUserId(userId));
        if (error) {
            dispatch({
                type: types.USER_PAGE_GET_URLS_ERROR,
                error,
            });
            throw error;
        }
        const { data: urls = [] } = resp;
        dispatch({
            type: types.USER_PAGE_GET_URLS_SUCCESS,
            urls,
        });
    },
};

const useUserPageActions = () => useActions(actions);

export { useUserPageActions };
export default reducer;
