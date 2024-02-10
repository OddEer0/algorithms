import { createInterface } from 'readline'; 
export const iojs = createInterface({ 
    input: process.stdin, 
    output: process.stdout 
});

export const scan = () => {
    return new Promise(resolve => {
        iojs.question("", input => {
            resolve(input); 
        });
    });
};
