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
<textarea class="okurigana2" onchange="updateSentence(this.value, ${index}, 'okuri2')"></textarea>
<textarea class="kaeriten" onchange="updateSentence(this.value, ${index}, 'kaeri')"></textarea>
<textarea class="okurigana" onchange="updateSentence(this.value, ${index}, 'okuri')"></textarea>
`;
        box.dataset.index = index; 
        container.appendChild(box);
    });}

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
    const line = document.getElementById('saidoku' + index);
    if (line.style.display == 'block') {
        line.style.display = 'none';
    } else {
        line.style.display = 'block';
    }
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
