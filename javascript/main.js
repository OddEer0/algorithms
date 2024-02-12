r = require("readline")
const iojs = r.createInterface({ 
    input: process.stdin, 
    output: process.stdout 
});

const scan = () => {
    return new Promise(resolve => {
        iojs.question("", input => {
            resolve(input); 
        });
    });
};

const main = async () => {
    const cases = await scan()
    const results = []

    for (let i = 0; i < +cases; i++) {
        const str = await scan()
        const maps = {
            t: 1,
            i: 1,
            n: 1,
            k: 1,
            o: 1,
            f: 2
        }

        for (let j = 0; j < str.length; j++) {
            if (maps[str[j].toLowerCase()] > 0) {
                maps[str[j].toLowerCase()]--
            } else {
                results[i] = "No"
                break
            }
        }
        if (results[i] !== "No") {
            results[i] = "Yes"
        }
    }

    for await (res of results) {
        console.log(res)
    }
    iojs.close();
}

main()