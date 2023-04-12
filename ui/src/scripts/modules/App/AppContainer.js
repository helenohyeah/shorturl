import React from 'react';
import { useSelector } from 'react-redux';
import Header from 'scripts/components/Header';
import { useAppActions } from './reducer/app';

const getState = state => state.app;

const AppContainer = ({
    // isPasswordHidden,
    // username,
    // password,
    children,
}) => {
    const userId = useSelector(state => getState(state).user.id);

    const { onLogout } = useAppActions();

    const navItems = !!userId
        ? [
              {
                  path: `/u/${userId}`,
                  label: 'My short urls',
              },
              {
                  path: '/logout',
                  label: 'Logout',
                  onClick: onLogout,
              },
          ]
        : [
              {
                  path: '/login',
                  label: 'Login',
              },
              {
                  path: '/register',
                  label: 'Register',
              },
          ];

    return (
        <div>
            <Header userId={userId} items={navItems} />
            <div>{children}</div>
        </div>
    );
};

export default AppContainer;
