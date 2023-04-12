import React from 'react';

const ShortenedUrl = ({ longUrl, shortUrl }) => {
    return (
        <div className="border rounded p-3">
            <p className="form-label">
                Long url:{' '}
                <a href={longUrl} target="_blank" rel="noreferrer">
                    {longUrl}
                </a>
            </p>
            <p className="form-label">
                Short url:{' '}
                <a href={shortUrl} target="_blank" rel="noreferrer">
                    {shortUrl}
                </a>
            </p>
        </div>
    );
};

export default ShortenedUrl;
