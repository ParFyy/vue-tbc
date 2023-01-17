<template>
    <div class="fixtures-container mx-auto">
        <v-card class="mb-1 mt-6 py-0 px-4 mt-sm-3 mb-sm-6 pa-sm-6 d-flex">
            <h2 class="mr-6 mr-sm-auto">{{ selected ? selected : $vuetify.display.xs ? "" : "All" }} Fixtures</h2>
            <div class="d-none d-sm-flex">
                <v-btn @click="selected='M1'">M1</v-btn>
                <v-btn @click="selected='M2'">M2</v-btn>
                <v-btn @click="selected='M3'">M3</v-btn>
                <v-btn @click="selected='L1'">L1</v-btn>
                <v-btn @click="selected='X1'">X1</v-btn>
                <v-btn @click="selected='X2'">M2</v-btn>
            </div>
            <v-select class="hidden-sm-and-up" v-model="selected" label="Select Team" :items="['M1', 'M2', 'M3', 'L1', 'X1', 'X2']" density="compact"></v-select>
        </v-card>
        <v-card class="ma-2 ma-sm-0">
            <v-table v-if="!$vuetify.display.xs || selected" density="default" hover="true">
                <thead>
                    <tr :class="{'small-table-font': $vuetify.display.xs}">
                        <th class="hidden-xs">Team</th>
                        <th>Opponent</th>
                        <th>Location</th>
                        <th>Date</th>
                        <th>Time</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="fixture in selectedFixtures" :key="fixture" :class="{'small-table-font': $vuetify.display.xs}">
                        <td class="hidden-xs">{{ fixture.Teamfor }}</td>
                        <td>{{ fixture.Teamagainst }}</td>
                        <td>{{ fixture.Location }}</td>
                        <td>{{ fixture.Date }}</td>
                        <td>{{ fixture.Time }}</td>
                    </tr>
                </tbody>
            </v-table>
        </v-card>
    </div>
</template>

<script>
import axios from "axios"
export default {
    data() {
        return {
            fixtures: undefined,
            selected: ""
        }
    },
    methods: {
        async getFixtures() {
            return (await axios.get("https://localhost:8080/GetTbcFixtures")).data;
        }
    },
    computed: {
        selectedFixtures() {
            var filtered = []
            if(this.fixtures) {
                this.fixtures.forEach(fixture => {
                    if (fixture.Teamfor === this.selected) {
                        filtered.push(fixture)
                    }
                });
                if (filtered.length == 0) {
                    return this.fixtures
                }
            }
            return filtered
        }
    },
    async created () {
        this.fixtures = await this.getFixtures();
    },
}
</script>

<style>
.fixtures-container{
    max-width: 700px;
}
.small-table-font td {
    font-size: 16px !important;
    padding-right: 6px !important;
    padding-left: 6px !important;
}
.small-table-font th {
    font-size: 16px !important;
    padding-right: 6px !important;
    padding-left: 6px !important;
}
</style>