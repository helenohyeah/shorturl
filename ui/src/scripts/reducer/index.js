import { configureStore } from '@reduxjs/toolkit';
import { combineReducers } from 'redux';
import app from 'scripts/modules/App/reducer/app';
import login from 'scripts/modules/LoginPage/reducer/loginPage';
import home from 'scripts/modules/HomePage/reducer/homePage';
import user from 'scripts/modules/UserPage/reducer/userPage';

function logger({ getState }) {
    return next => action => {
        console.log('ACTION: ', action);

        // Call the next dispatch method in the middleware chain.
        const returnValue = next(action);

        console.log('STATE: ', getState());

        // This will likely be the action itself, unless
        // a middleware further in chain changed it.
        return returnValue;
    };
}

const store = configureStore({
    middleware: getDefaultMiddleware =>
        getDefaultMiddleware({ serializableCheck: false }).concat(logger),
    reducer: combineReducers({
        login,
        app,
        home,
        user,
    }),
});

export default store;
