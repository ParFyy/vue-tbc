<template>
    <div class="fixtures-container mx-auto">
        <v-card class="mb-1 mt-6 py-0 px-4 mt-sm-3 mb-sm-6 pa-sm-6 d-flex">
            <h2 class="mr-6 mr-sm-auto">All Fixtures</h2>
            <div class="d-none d-sm-flex">
                <v-btn>M1</v-btn>
                <v-btn>M2</v-btn>
                <v-btn>M3</v-btn>
                <v-btn>L1</v-btn>
                <v-btn>X1</v-btn>
                <v-btn>M2</v-btn>
            </div>
            <v-select class="hidden-sm-and-up" label="Select Team" :items="['M1', 'M2', 'M3', 'L1', 'X1', 'X2']" density="compact"></v-select>
        </v-card>
        <v-card>
            <v-table density="default" hover="true">
                <thead>
                    <tr :class="{'small-table-font': $vuetify.display.xs}">
                        <th>Team</th>
                        <th>Opponent</th>
                        <th>Location</th>
                        <th>Date</th>
                        <th>Time</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="fixture in fixtures" :key="fixture" :class="{'small-table-font': $vuetify.display.xs}">
                        <td>{{ fixture.Teamfor }}</td>
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
            fixtures: undefined
        }
    },
    methods: {
        async getFixtures() {
            return (await axios.get("http://localhost:8080/GetTbcFixtures")).data;
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
    padding-right: 5px !important;
    padding-left: 5px !important;
    text-align: center;
}
.small-table-font th {
    font-size: 16px !important;
    padding-right: 5px !important;
    padding-left: 5px !important;
    text-align: center !important;
}
</style>