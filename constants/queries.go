package constants

const (
	QueryDeposit = `
		INSERT INTO ledgers (user_id, amount, balance, type, created_at)
		VALUES ($1, $2, (SELECT COALESCE(MAX(balance), '0')::numeric + $2 FROM ledgers WHERE user_id = $1), 'deposit', $3)
	`

	QueryWithdraw = `INSERT INTO ledgers (user_id, amount, balance, type, created_at)
		SELECT $1, -CAST($2 AS DECIMAL(36,18)), balance - CAST($2 AS DECIMAL(36,18)), 'withdraw', $3
		FROM ledgers 
		WHERE user_id = $1 AND balance >= CAST($2 AS DECIMAL(36,18))
		ORDER BY created_at DESC 
		LIMIT 1
	`

	QueryTransferOut = `INSERT INTO ledgers (user_id, amount, balance, type, created_at)
			SELECT $1, -CAST($2 AS DECIMAL(36,18)), balance - CAST($2 AS DECIMAL(36,18)), 'transfer_out', $3
			FROM ledgers 
			WHERE user_id = $1 
			AND balance >= CAST($2 AS DECIMAL(36,18))
			ORDER BY created_at DESC 
			LIMIT 1`

	QueryTransferIn = `INSERT INTO ledgers (user_id, amount, balance, type, created_at)
			VALUES (
				$1, 
				$2, 
				COALESCE(
					(SELECT balance FROM ledgers WHERE user_id = $1 ORDER BY created_at DESC LIMIT 1), 
					0
				) + $2,
				'transfer_in', 
				$3
			)`

	QueryWalletBalance = `SELECT COALESCE((SELECT balance FROM ledgers WHERE user_id = $1 ORDER BY created_at DESC LIMIT 1), 0)`

	QueryWalletHistory = `SELECT amount, balance, type, created_at FROM ledgers WHERE user_id = $1 ORDER BY created_at DESC`
)
