import React from 'react';
import ReactDOM from 'react-dom';
import { Provider } from 'react-redux';
import Router from 'scripts/router/router';
import store from 'scripts/reducer';

import 'styles/_reset.scss';
import '@coreui/coreui/dist/css/coreui.min.css';

ReactDOM.render(
    <React.StrictMode>
        <Provider store={store}>
            <Router />
        </Provider>
    </React.StrictMode>,
    document.getElementById('root')
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
