import React, { useState } from 'react';
import { useSelector } from 'react-redux';
import { of } from 'await-of';
import { useHomePageActions } from './reducer/homePage';
import * as API from 'request';
import HomePage from './components/HomePage';

const getState = state => state.home;

const HomePageContainer = () => {
    const { onFormChange, onShortenUrl, onResetUI } = useHomePageActions();

    const isShortened = useSelector(state => getState(state).isShortened);
    const longUrl = useSelector(state => getState(state).longUrl);
    const shortUrl = useSelector(state => getState(state).shortUrl);
    const userId = useSelector(state => state.app.user.id);

    const [error, setError] = useState();

    const _onChange = name => event => {
        const value = event.target.value;
        onFormChange(name, value);
    };

    const _onSubmit = async event => {
        event.preventDefault();
        event.stopPropagation();

        // todo: validate longUrl looks like a website
        const [resp = {}, err] = await of(
            API.shortenUrl({
                userId,
                originalUrl: longUrl,
            })
        );

        if (err) {
            setError(err.message);
            throw err;
        }
        setError();
        const { data: url = {} } = resp;
        onShortenUrl(url.shortUrl, url.redirectUrl);
    };

    return (
        <HomePage
            isShortened={isShortened}
            longUrl={longUrl}
            shortUrl={shortUrl}
            error={error}
            onChange={_onChange}
            onSubmit={_onSubmit}
            onShortenAgain={onResetUI}
        />
    );
};

export default HomePageContainer;
