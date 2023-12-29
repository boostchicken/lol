

type lolEntry = {
    Command: string,
    Value: string
}

let entries: lolEntry[] = [{Command: "github", Value: "League of Legends"}];
const config= fetch("https://lol.boostchicken.io/liveconfig")
.then((res) => res.json()).then((json) => {entries.push(json); console.log(json);});

function xml_escape(unsafe: string) {
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
    const results: chrome.omnibox.SuggestResult[] = []
    const action = text.trim();
    const args: string[] = action.split(' ');
    console.log(args);
    for (let elm of entries) {
        if (elm.Command.includes(args[0])) {
            results.push({
                content: elm.Command,
                description: `${elm.Command}<dim> - ${xml_escape(elm.Value)}</dim>`
            })
        }
    }
    suggest(results)
});

chrome.omnibox.onInputEntered.addListener((term, OnInputEnteredDisposition) => {
    let tab_disposition: chrome.search.Disposition = "NEW_TAB";
    switch (OnInputEnteredDisposition) {
        case "currentTab":
            tab_disposition = "CURRENT_TAB";
            break;
    }
    chrome.search.query({
        disposition: tab_disposition,
        text: term
    });

});