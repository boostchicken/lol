let conf = {
}

let url = "https://lol.boostchicken.dev/liveconfig"
async function cache() {
    try {
        fetch(url)
        .then((response) => response.json())
        .then((data) => conf = data)
    } catch (e) {
        console.log(e)
    } finally {
        setInterval(cache,60000)
    }
}   

function xml_escape(unsafe) {
    return unsafe.replace(/[<>&'"]/g, function (c) {
        switch (c) {
            case '<': return '&lt;';
            case '>': return '&gt;';
            case '&': return '&amp;';
            case '\'': return '&apos;';
            case '"': return '&quot;';
        }
    });
}

chrome.omnibox.onInputChanged.addListener((text, suggest) => {
    const results = []
    const action = text.trim();
    const args = action.split(' ');
    if(args[0] == "setUrl") {
        url = args[1];
        results.push({content: "setUrl", description: `setUrl ${url} for liveconfig (e.g setUrl https://lol.boostchicken.dev/liveconfig)`})
        suggest(results)
    }
    for (const elm of conf.Entries) {
        if(elm.Command.includes(args[0])) {
            results.push({
                content: elm.Command,
                description: `${elm.Command}<dim> - ${xml_escape(elm.Value)}</dim>`
            })
        }
    }
    suggest(results)
});

chrome.omnibox.onInputEntered.addListener((term, OnInputEnteredDisposition) => {
    let tab_disposition = ""
    switch (OnInputEnteredDisposition) {
        case "currentTab":
            tab_disposition = "CURRENT_TAB";
            break;
        default:
            tab_disposition = "NEW_TAB";
            break;
    }
    if(term.split(" ")[0] == "setUrl") {
        url = term.split(" ")[1]
    } else {
        chrome.search.query({
            disposition: tab_disposition, 
            text: term
        });
    }
});

cache()