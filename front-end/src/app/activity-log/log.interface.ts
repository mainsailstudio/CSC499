export interface ConfigActivity {
    userID: number;
    totalCreationTime: number;
    avgSecretLength: number;
}

export interface LoginActivity {
    userID: number;
    testLevel: number;
    loginTime: number;
    failures: number;
    refreshes: number;
    secretLength: number;
}
