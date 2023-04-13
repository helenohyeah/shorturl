import React from 'react';
import { useSelector } from 'react-redux';
import { of } from 'await-of';
import { useHomePageActions } from './reducer/homePage';
import HomePage from './components/HomePage';

const getState = state => state.home;

const HomePageContainer = () => {
    const { onFormChange, onShortenUrl, onResetUI } = useHomePageActions();

    const isShortened = useSelector(state => getState(state).isShortened);
    const longUrl = useSelector(state => getState(state).longUrl);
    const shortUrl = useSelector(state => getState(state).shortUrl);
    const error = useSelector(state => getState(state).error);
    const userId = useSelector(state => state.app.user.id);

    const _onChange = name => event => {
        const value = event.target.value;
        onFormChange(name, value);
    };

    const _onSubmit = async event => {
        event.preventDefault();
        event.stopPropagation();

        // todo: validate longUrl looks like a website before sending req

        const [, err] = await of(onShortenUrl(userId, longUrl));

        if (err) {
            throw err;
        }
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
