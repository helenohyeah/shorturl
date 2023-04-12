import { configureStore } from '@reduxjs/toolkit';
import { combineReducers } from 'redux';
import login from 'scripts/modules/LoginPage/reducer/loginPage';

const store = configureStore({
    middleware: (
        getDefaultMiddleware
        // NOTE: if you need to disable the immutability check - update the line below to:
        // ) => getDefaultMiddleware({ serializableCheck: false, immutableCheck: false }).concat(...middlewares),
    ) => getDefaultMiddleware({ serializableCheck: false }),
    reducer: combineReducers({
        login,
    }),
});

export default store;
