import React from 'react';
import { BrowserRouter as Router, Switch, Route, Redirect } from 'react-router-dom';
import App from 'scripts/modules/App';
import HomePage from 'scripts/modules/HomePage';
import LoginPage from 'scripts/modules/LoginPage';
import UserPage from 'scripts/modules/UserPage';

const AppRouter = () => {
    return (
        <Router>
            <Switch>
                <Route
                    exact
                    path="/"
                    render={() => (
                        <App>
                            <HomePage />
                        </App>
                    )}
                />
                <Route
                    exact
                    path="/login"
                    render={() => (
                        <App>
                            <LoginPage />
                        </App>
                    )}
                />
                <Route exact path="/logout" render={() => <Redirect />} />
                <Route
                    exact
                    path="/register"
                    render={() => (
                        <App>
                            <LoginPage isRegister />
                        </App>
                    )}
                />
                <Route
                    path="/u/:userId"
                    render={() => (
                        <App>
                            <UserPage />
                        </App>
                    )}
                />
            </Switch>
        </Router>
    );
};

export default AppRouter;
