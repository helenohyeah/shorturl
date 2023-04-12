import React, { useEffect } from 'react';
import { shallowEqual, useSelector } from 'react-redux';
import { useParams } from 'react-router-dom';
import { useUserPageActions } from './reducer/userPage';
import UserPage from './components/UserPage';

const getState = state => state.user;

const UserPageContainer = () => {
    const { userId } = useParams();
    const { onResetUI, onGetUrls } = useUserPageActions();

    const isFetched = useSelector(state => getState(state).isFetched);
    const urls = useSelector(state => getState(state).urls, shallowEqual);

    useEffect(() => {
        if (!userId) {
            return;
        }
        onGetUrls(userId);
    }, [onGetUrls, userId]);

    useEffect(() => {
        if (isFetched || !userId) {
            return;
        }
        onGetUrls(userId);
        return () => {
            onResetUI();
        };
    }, [onGetUrls, onResetUI, isFetched, userId]);

    return <UserPage urls={urls} />;
};

export default UserPageContainer;
