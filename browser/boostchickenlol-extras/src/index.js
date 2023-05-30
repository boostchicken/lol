let conf = {}

function cache() {
    fetch("https://lol.boostchicken.dev/liveconfig")
    .then((response) => response.json())
    .then((data) => conf = data)
    setInterval(cache,60000)
}   

function xml_escape (text) {
    return text.replace(/[<>&'"]/g, ch => {
        switch (ch) {
        case '<': return '&lt;'
        case '>': return '&gt;'
        case '&': return '&amp;'
        case '\'': return '&apos;'
        case '"': return '&quot;'
        }
    })
}

chrome.omnibox.onInputChanged.addListener((text, suggest) => {
    const results =[]
    for (const elm of conf.Entries) {
        if(elm.Command.includes(text)) {
            results.push({
                    content: text,
                    deletable: false,
                    description: `<match>${xml_escape(elm.Command)}</match>`
                    + '<dim>' + xml_escape(elm.Value) + '</dim>'
                })
        }
    }
    console.log(results)
    suggest(results)
});

chrome.omnibox.setDefaultSuggestion({description:  "Search LoL Command"})


chrome.omnibox.onInputEntered.addListener((text, OnInputEnteredDisposition) => {
        if (OnInputEnteredDisposition === 'currentTab') {
            chrome.tabs.create(text);
        } else {
            chrome.tabs.update(text);
        }
    });

cache()