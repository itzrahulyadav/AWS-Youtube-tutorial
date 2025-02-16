import * as fs from 'node:fs';

const html = fs.readFileSync('index.html', { encoding: 'utf8' });

export const handler = async (event) => {    
    const response = {
        statusCode: 200,
        headers: {
            'Content-Type': 'text/html',
        },
        body: html,
    };
    return response;
};
