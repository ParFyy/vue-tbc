<template>
    <table class="table">
        <thead>
            <tr>
                <th>Team</th>
                <th>Opponent</th>
                <th>Location</th>
                <th>Date</th>
                <th>Start Time</th>
            </tr>
        </thead>
    </table>
</template>

<script>
    async function getFixtures(){
        var fixtures = await (await fetch("https://agile-waters-57617.herokuapp.com/https://ndblfixturesoftware.herokuapp.com/GetTbcFixtures")).json()
        var table = document.getElementsByClassName("table")[0]
        fixtures.forEach(fixture => {
            var tableRow = ""
            tableRow += "<tr>"
            tableRow += "<td>" + fixture.Teamfor + "</td>"
            tableRow += "<td>" + fixture.Teamagainst + "</td>"
            tableRow += "<td>" + fixture.Location + "</td>"
            tableRow += "<td>" + fixture.Date + "</td>"
            tableRow += "<td>" + fixture.Time + "</td>"
            tableRow += "</tr>"
            table.innerHTML += tableRow
        });
    }
    getFixtures()
</script>

<style scoped>
    table {
        max-width: fit-content;
    }
</style>