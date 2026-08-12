const rawInput = document.getElementById('input');
const inputButton = document.getElementById('btn-input');
const tategaki = document.querySelector('.tategaki');
const kakikudashi = document.querySelector('.kakikudashi');
const footer = document.querySelector('.footer');

async function renderTategaki(sentence) {
    const container = document.getElementById("tategaki");
    container.innerHTML = "";

    characters = sentence.Characters
    characters.forEach((char, index) => {
        const box = document.createElement("div");
        box.className = "kanji-box";
        box.innerHTML = `
<span onclick="charLookup(this)">${char.Kanji}</span>
<button class="btn-issaidoku" onclick="updateSaidoku(${index})">再</button>
<div class="line-s" id="saidoku${index}"></div>
<button class="btn-isjuku" onclick="updateJuku(${index})">熟</button>
<div class="line-j" id="juku${index}"></div>
<textarea class="kaeriten" onchange="updateSentence(this.value, ${index}, 'kaeri')"></textarea>
<textarea class="okurigana" onchange="updateSentence(this.value, ${index}, 'okuri')"></textarea>
`;
        box.dataset.index = index;
        container.appendChild(box);
    });
}

inputButton.addEventListener('click', async () => {
    const text = rawInput.value;
    if (!text) return;
    kakikudashi.innerHTML = text;

    try {
        const response = await fetch('/api/sentences', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ text: text })
        });

        if (!response.ok) {
            throw new Error('Response failed: ' + response.statusText);
        }

        const sentenceData = await response.json();
        renderTategaki(sentenceData.sentence);

    } catch (error) {
        console.error('Error:', error);
        alert("Failure; check the console");
    }
});

function updateSaidoku(index) {
    updateSentence('', index, 'saidoku');
    const box = document.querySelector(`.kanji-box[data-index="${index}"]`);
    let textarea = box.querySelector('.okurigana2');
    const line = box.querySelector('.line-s');

    if (!textarea) {
        textarea = document.createElement("textarea");
        textarea.className = "okurigana2";
        textarea.onchange = function() { updateSentence(this.value, index, 'okuri2'); };
        box.appendChild(textarea); // Don't forget to append the textarea to the box!
    }

    // Get the computed style
    const textareaDisplay = window.getComputedStyle(textarea).display;
    const lineDisplay = window.getComputedStyle(line).display;

    // Toggle visibility
    textarea.style.display = textareaDisplay === 'block' ? 'none' : 'block';
    line.style.display = lineDisplay === 'block' ? 'none' : 'block';
}

function updateJuku(index) {
    updateSentence('', index, 'juku');
    const line = document.getElementById('juku' + index);
    if (line.style.display == 'block') {
        line.style.display = 'none';
    } else {
        line.style.display = 'block';
    }
}

async function updateSentence(value, index, type) {
    try {
        const response = await fetch('/api/sentences/1', {
            method: 'PATCH',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ index: index, text: value, sentence_id: 1 , type: type})
        });

        if (!response.ok) {
            throw new Error('Response failed: ' + response.statusText);
        }

        const kakikudashibun = await response.json();
        kakikudashi.innerHTML = kakikudashibun.text

    } catch (error) {
        console.error('Error:', error);
        alert("Failure; check the console");
    }
};

async function charLookup(element) {
    try {
        const response = await fetch('/api/characters/'+element.textContent, {
            method: 'GET'
        });

        if (!response.ok) {
            throw new Error('Response failed: ' + response.statusText);
        }

        const lookupResponse = await response.json();
        const imi = lookupResponse.imi.replace(/\n/g, '<br>');
        footer.innerHTML = imi;

    } catch (error) {
        console.error('Error:', error);
        alert("Failure; check the console");
    }
}
