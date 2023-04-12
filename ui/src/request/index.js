import request from 'superagent';
import { format } from 'util';

const BASE_URL = 'http://localhost:8080';

const API_ENDPOINTS = {
    shortenUrl: 'shorten_url',
    register: 'register',
    login: 'acct_login',
    userUrls: 'urls_user/%s',
};

const getAPIUrl = path => {
    return `${BASE_URL}/${path}`;
};

const apiCall = args => {
    const { method, url, data } = args;
    const methodType = method.toLowerCase();
    const requestMethod = request[methodType];

    return new Promise((resolve, reject) => {
        requestMethod(url)
            .set('Content-Type', 'application/json')
            .set('Accept', 'application/json')
            .send(data)
            .end((err, res) => {
                if (err) {
                    // extract error message in response body
                    return reject({
                        error: err,
                        message: err.response?.body?.error,
                    });
                }
                // extract data in body
                resolve(res.body);
            });
    });
};

export const register = data => {
    return apiCall({
        method: 'post',
        url: getAPIUrl(API_ENDPOINTS.register),
        data,
    });
};

export const login = data => {
    return apiCall({
        method: 'post',
        url: getAPIUrl(API_ENDPOINTS.login),
        data,
    });
};

export const shortenUrl = data => {
    return apiCall({
        method: 'post',
        url: getAPIUrl(API_ENDPOINTS.shortenUrl),
        data,
    });
};

export const getUrlsByUserId = userId => {
    return apiCall({
        method: 'get',
        url: format(getAPIUrl(API_ENDPOINTS.userUrls), userId),
    });
};
