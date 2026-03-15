const rawInput = document.getElementById('input');
const inputButton = document.getElementById('btn-input');
const tategaki = document.querySelector('.tategaki');
const kakikudashi = document.querySelector('.kakikudashi');

async function renderTategaki(sentence) {
    const container = document.getElementById("tategaki");
    container.innerHTML = ""; 

    characters = sentence.Characters
    characters.forEach((char, index) => {
        const box = document.createElement("div");
        box.className = "kanji-box";
        box.innerHTML = `
<span>${char.Kanji}</span>
<textarea class="kaeriten" onchange="updateKaeriten(this.value, ${index})"></textarea>
<!-- <textarea class="okurigana2"></textarea> -->
<textarea class="okurigana" onchange="updateOkuri(this.value, ${index})"></textarea>
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

async function updateKaeriten(value, index) {
    try {
        const response = await fetch('/api/sentences/1', {
            method: 'PATCH',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ index: index, kaeriten: value, sentence_id: 1, update_type: "kaeri"})
        });

        if (!response.ok) {
            throw new Error('Response failed: ' + response.statusText);
        }

        const kakikudashibun = await response.json();
        kakikudashi.innerHTML = kakikudashibun.text 

    } catch (error) {
        console.error('Erorr:', error);
            alert("Failure; check the console");
    }
};

async function updateOkuri(value, index) {
    try {
        const response = await fetch('/api/sentences/1', {
            method: 'PATCH',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ index: index, okuri: value, sentenceId: 1 , update_type:"okuri"})
        });

        if (!response.ok) {
            throw new Error('Response failed: ' + response.statusText);
        }

        const kakikudashibun = await response.json();
        kakikudashi.innerHTML = kakikudashibun.text 

    } catch (error) {
        console.error('Erorr:', error);
            alert("Failure; check the console");
    }
};
